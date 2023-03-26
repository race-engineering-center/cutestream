#ifndef GENERATOR_H
#define GENERATOR_H

#include <QList>
#include <QDate>
#include <QUuid>
#include <QTime>
#include <QDateTime>

/*
 *
 * This file contains template function for generating
 * test data lists
 *
 */

template <typename T>
QList<T> getTestData();

template<std::integral T>
QList<T> getTestData() {
    QList<T> result;
    if constexpr (std::is_unsigned_v<T>) {
        result = {
            0,
            42,
            std::numeric_limits<T>::max()
        };
    }
    else {
        result = {
            0,
            42,
            -73,
            std::numeric_limits<T>::min(),
            std::numeric_limits<T>::max()
        };
    }
    return result;
}

template<std::floating_point T>
QList<T> getTestData() {
    return {
        static_cast<T>(0),
                static_cast<T>(3.1415),
                static_cast<T>(-9000),
                static_cast<T>(2.71828)
    };
}




#endif // GENERATOR_H
