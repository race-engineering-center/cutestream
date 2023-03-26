#include "representator.h"

template<>
QJsonValue getRepresentation(const QUuid& value) {
    QJsonValue result = value.toString(QUuid::Id128);
    return result;
}

template<>
QJsonValue getRepresentation(const QDate& value) {
    QJsonObject result;
    result["year"] = value.year();
    result["month"] = value.month();
    result["day"] = value.day();
    return result;
}

template<>
QJsonValue getRepresentation(const QTime& value) {
    QJsonObject result;

    result["hour"] = value.hour();
    result["minute"] = value.minute();
    result["sec"] = value.second();
    result["ms"] = value.msec();

    return result;
}

template<>
QJsonValue getRepresentation(const QDateTime& value) {
    QJsonObject result;
    result["year"] = value.date().year();
    result["month"] = value.date().month();
    result["day"] = value.date().day();

    result["hour"] = value.time().hour();
    result["minute"] = value.time().minute();
    result["sec"] = value.time().second();
    result["ms"] = value.time().msec();

    return result;
}
