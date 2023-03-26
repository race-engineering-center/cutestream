#include "qdatastreamwriter.h"

QDataStreamWriter::QDataStreamWriter() :
    m_precision(QDataStream::DoublePrecision)
{

}

QDataStream::FloatingPointPrecision QDataStreamWriter::precision() const
{
    return m_precision;
}

void QDataStreamWriter::setPrecision(QDataStream::FloatingPointPrecision newPrecision)
{
    m_precision = newPrecision;
}
