// Code generated by Microsoft Kiota - DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package security
type SensorHealthStatus int

const (
    HEALTHY_SENSORHEALTHSTATUS SensorHealthStatus = iota
    NOTHEALTHYLOW_SENSORHEALTHSTATUS
    NOTHEALTHYMEDIUM_SENSORHEALTHSTATUS
    NOTHEALTHYHIGH_SENSORHEALTHSTATUS
    UNKNOWNFUTUREVALUE_SENSORHEALTHSTATUS
)

func (i SensorHealthStatus) String() string {
    return []string{"healthy", "notHealthyLow", "notHealthyMedium", "notHealthyHigh", "unknownFutureValue"}[i]
}
func ParseSensorHealthStatus(v string) (any, error) {
    result := HEALTHY_SENSORHEALTHSTATUS
    switch v {
        case "healthy":
            result = HEALTHY_SENSORHEALTHSTATUS
        case "notHealthyLow":
            result = NOTHEALTHYLOW_SENSORHEALTHSTATUS
        case "notHealthyMedium":
            result = NOTHEALTHYMEDIUM_SENSORHEALTHSTATUS
        case "notHealthyHigh":
            result = NOTHEALTHYHIGH_SENSORHEALTHSTATUS
        case "unknownFutureValue":
            result = UNKNOWNFUTUREVALUE_SENSORHEALTHSTATUS
        default:
            return nil, nil
    }
    return &result, nil
}
func SerializeSensorHealthStatus(values []SensorHealthStatus) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i SensorHealthStatus) isMultiValue() bool {
    return false
}
