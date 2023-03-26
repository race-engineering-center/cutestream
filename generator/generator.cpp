#include "generator.h"

template<>
QList<QUuid> getTestData() {
    return {
        QUuid::createUuid(),
                QUuid::createUuid(),
                QUuid::createUuid(),
                QUuid::createUuid(),
                QUuid::createUuid(),
    };
}

template<>
QList<QDate> getTestData() {
    return {
        QDate(1998, 07, 25),
                QDate(1995, 05, 20),
                QDate(2022, 05, 03),
                QDate(2022, 06, 17)
    };
}

template<>
QList<QTime> getTestData() {
    return {
        QTime(0, 0, 0),
                QTime(2, 42, 31, 123),
                QTime(12, 0, 30, 250),
                QTime(16, 45, 0),
    };
}

template<>
QList<QDateTime> getTestData() {
    return {
        QDateTime(QDate(1998, 07, 25), QTime(0, 0, 0)),
                QDateTime(QDate(1995, 05, 20), QTime(2, 42, 31, 123)),
                QDateTime(QDate(2022, 05, 03), QTime(12, 0, 30, 250)),
                QDateTime(QDate(2022, 06, 1), QTime(16, 45, 0))
    };
}
