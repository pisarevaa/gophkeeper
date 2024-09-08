package model

import (
	"time"
)

type GeoLog struct {
	Logs []GeoLogOne `json:"logs" validate:"required"` // список логов
}

type GeoLogOne struct {
	Date      time.Time     `json:"date"                validate:"required" example:"2019-04-01T00:00:00+01:00"` // время события
	Longitude float64       `json:"longitude,omitempty"`                                                         // долгота (координаты события)
	Latitude  float64       `json:"latitude,omitempty"`                                                          // широта (координаты события)
	Memory    GeoLogMemory  `json:"memory,omitempty"`                                                            // параметры памяти
	Battery   GeoLogBattery `json:"battery,omitempty"`                                                           // параметры батареи
	Network   GeoLogNetwork `json:"network,omitempty"`                                                           // параметры сети

}

type GeoLogMemory struct {
	FreeRAM string `json:"freeRam,omitempty"` // Свободная ОЗУ (пример, 848,35 МB (32,21%))
	FreeHdd string `json:"freeHdd,omitempty"` // Свободно внутренней памяти (пример, 1,17 Gb (5,91 %))
}

type GeoLogBattery struct {
	BatteryChargeLevel string `json:"batteryChargeLevel,omitempty"` // Уровень заряда батареи (пример, 25,0 %)
	Temperature        string `json:"temperature,omitempty"`        // Температура (пример, 32,8 ℃ (91,0 ℉))
	BatteryCapacity    string `json:"batteryCapacity,omitempty"`    // Емкость батареи (пример, 4000 mAh)
	BatteryMode        string `json:"batteryMode,omitempty"`        // Режим батареи (пример, Зарядка)
	BatteryStatus      string `json:"batteryStatus,omitempty"`      // Состояние батареи (пример, Хорошо)
	PowerSource        string `json:"powerSource,omitempty"`        // Источник питания (пример, Аккумулятор)
}

type GeoLogNetwork struct {
	NetworkDataType  string `json:"networkDataType,omitempty"`  // Тип данных (пример, WiFi)
	Connection       string `json:"connection,omitempty"`       // Подключение для передачи данных (пример, Подключен)
	NetworkName      string `json:"networkName,omitempty"`      // Наименование сети (пример, Домашняя)
	ConnectionSpeed  string `json:"connectionSpeed,omitempty"`  // Скорость соединения (пример, 90 Мб/с)
	WifiTraffic      string `json:"wifiTraffic,omitempty"`      // Потребленный трафик WiFi за период (пример, 1,58 Gb)
	MobileTraffic    string `json:"mobileTraffic,omitempty"`    // Потребленный трафик мобильный (пример, 25,1 Gb)
	SignalStrength   string `json:"signalStrength,omitempty"`   // Мощность WiFi (пример, значения: 0.0 - нет сигнала/слабый сигнал 1.0 - хороший сигнал)
	MobileOperator   string `json:"mobileOperator,omitempty"`   // Название сотового оператора (пример, A1)
	Gprs             string `json:"gprs,omitempty"`             // Наличие GPRS (пример, 3G или 4G)
	CallID           string `json:"callId,omitempty"`           // Call ID устройства (пример, 123456789)
	Imei             string `json:"imei,omitempty"`             // IMEI код устройства (пример, 35-419002-389644-3)
	SnSim            string `json:"snSim,omitempty"`            // Серийный номер SIM карты (пример, 89 7 0199)
	SignalLevel      string `json:"signalLevel,omitempty"`      // Уровень сигнала в dBm (пример, 2 dBm)
	DataTransferRate string `json:"dataTransferRate,omitempty"` // Скорость передачи данных (пример, 100 mbs)
	IsVpn            bool   `json:"isVpn,omitempty"`            // Подключен ли к VPN
}
