#ifndef QDATASTREAMWRITER_H
#define QDATASTREAMWRITER_H

#include <stdexcept>

#include <QDataStream>
#include <QJsonObject>
#include <QJsonArray>

#include "generator.h"
#include "representator.h"

class QDataStreamWriter
{
public:
    QDataStreamWriter();

    template <typename T>
    QString getBase64(const T& value, int version) {
        QByteArray bytes;
        QDataStream stream(&bytes, QIODeviceBase::WriteOnly);
        stream.setFloatingPointPrecision(m_precision);
        stream.setVersion(version);
        stream << value;
        return bytes.toBase64();
    }

    template <typename T>
    QJsonArray getJson(int version) {
        QJsonArray result;
        auto values = getTestData<T>();
        for (const auto& v: values) {
            QJsonObject object;
            object["value"] = getRepresentation(v);
            object["serialized"] = getBase64(v, version);
            result.append(object);
        }
        return result;
    }

    QDataStream::FloatingPointPrecision precision() const;
    void setPrecision(QDataStream::FloatingPointPrecision p);

private:
    QDataStream::FloatingPointPrecision m_precision;
};

#endif // QDATASTREAMWRITER_H
