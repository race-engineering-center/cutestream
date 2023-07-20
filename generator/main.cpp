#include <QCoreApplication>
#include <QJsonObject>
#include <QJsonArray>
#include <QJsonDocument>
#include <QFile>
#include <QDir>
#include <QDebug>
#include <QProcessEnvironment>
#include <QUuid>

#include "qdatastreamwriter.h"

template <typename Callable>
void generate(QString fileName, QDir rootDir, const QVector<int>& versions, Callable func)
{
    QJsonObject output;
    for (auto version: versions) {
        QJsonObject versionObject = func(version);
        output[QString::number(version)] = versionObject;
    }
    QFile file(rootDir.filePath(fileName));
    if (!file.open(QIODeviceBase::WriteOnly)) {
        throw std::invalid_argument(QString("Unable to write to " + file.fileName()).toStdString());
    }

    file.write(QJsonDocument(output).toJson());
}

int main(int argc, char *argv[]) {
    QVector<int> versions {
        QDataStream::Qt_5_13,
                QDataStream::Qt_6_0
    };

    QJsonObject output;

    QString rootPath = QProcessEnvironment::systemEnvironment().value("CUTESTREAM_TEST_DIR");
    QDir rootDir(rootPath);
    if (rootPath.isEmpty() || !rootDir.exists()) {
        qDebug()<<"CUTESTREAM_TEST_DIR environment variable not set or points to incorrect folder";
        return 1;
    }

    generate("generated_int.json", rootDir, versions, [](int version) -> QJsonObject {
        QDataStreamWriter writer;
        QJsonObject versionObject;

        versionObject["int8"] = writer.getJson<int8_t>(version);
        versionObject["uint8"] = writer.getJson<uint8_t>(version);

        versionObject["int16"] = writer.getJson<int16_t>(version);
        versionObject["uint16"] = writer.getJson<uint16_t>(version);

        versionObject["int32"] = writer.getJson<int32_t>(version);
        versionObject["uint32"] = writer.getJson<uint32_t>(version);

        versionObject["int64"] = writer.getJson<int64_t>(version);
        versionObject["uint64"] = writer.getJson<uint64_t>(version);

        return versionObject;
    });

    generate("generated_float.json", rootDir, versions, [](int version) -> QJsonObject {
        QDataStreamWriter writer;
        QJsonObject versionObject;

        writer.setPrecision(QDataStream::DoublePrecision);
        versionObject["float_d"] = writer.getJson<float>(version);
        versionObject["double_d"] = writer.getJson<double>(version);

        writer.setPrecision(QDataStream::SinglePrecision);
        versionObject["float_s"] = writer.getJson<float>(version);
        versionObject["double_s"] = writer.getJson<double>(version);

        return versionObject;
    });

    generate("generated_uuid.json", rootDir, versions, [](int version) -> QJsonObject {
        QDataStreamWriter writer;
        QJsonObject versionObject;
        versionObject["uuid"] = writer.getJson<QUuid>(version);
        return versionObject;
    });

    generate("generated_date.json", rootDir, versions, [](int version) -> QJsonObject {
        QDataStreamWriter writer;
        QJsonObject versionObject;
        versionObject["date"] = writer.getJson<QDate>(version);
        return versionObject;
    });

    generate("generated_time.json", rootDir, versions, [](int version) -> QJsonObject {
        QDataStreamWriter writer;
        QJsonObject versionObject;
        versionObject["time"] = writer.getJson<QTime>(version);
        return versionObject;
    });

    generate("generated_datetime.json", rootDir, versions, [](int version) -> QJsonObject {
        QDataStreamWriter writer;
        QJsonObject versionObject;
        versionObject["datetime"] = writer.getJson<QDateTime>(version);
        return versionObject;
    });

    return 0;
}
