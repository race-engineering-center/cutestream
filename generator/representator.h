#ifndef REPRESENTATOR_H
#define REPRESENTATOR_H

#include <QJsonObject>

template<typename T>
QJsonValue getRepresentation(const T& value);

template<typename T>
concept Numeric = std::integral<T> || std::floating_point<T>;

template<Numeric T>
inline QJsonValue getRepresentation(const T &value) {
    QJsonValue result = QString::number(value);
    return result;
}

#endif // REPRESENTATOR_H
