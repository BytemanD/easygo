package ini

import (
	"github.com/fjboy/magic-pocket/pkg/global/logging"
	"github.com/go-ini/ini"
)

var CONF Conf

type Conf struct {
	file    string
	iniFile ini.File
}

func (conf *Conf) SetBlockMode(blockMode bool) {
	conf.iniFile.BlockMode = blockMode
}

func (conf *Conf) Get(section string, key string) string {
	if !conf.iniFile.Section(section).HasKey(key) {
		return ""
	}
	return conf.iniFile.Section(section).Key(key).Value()
}

func (conf *Conf) GetDefault(key string) string {
	return conf.Get("", key)
}

func (conf *Conf) Set(section string, key string, value string) {
	conf.iniFile.Section(section).Key(key).SetValue(value)
	conf.iniFile.SaveTo(conf.file)
}
func (conf *Conf) Delete(section string, key string) {
	if !conf.iniFile.Section(section).HasKey(key) {
		return
	}
	conf.iniFile.Section(section).DeleteKey(key)
	conf.iniFile.SaveTo(conf.file)
}

func (conf *Conf) Log(info bool) {
	logFunc := logging.Debug
	if info {
		logFunc = logging.Info
	}
	for _, section := range conf.iniFile.Sections() {
		for _, key := range section.Keys() {
			logFunc("config: %s.%s = %s",
				section.Name(), key.Name(), key.Value())
		}
	}
}

func (conf *Conf) Load(source string) error {
	iniFile, err := ini.Load(source)
	if err != nil {
		return err
	}
	conf.file = source
	conf.iniFile = *iniFile
	return nil
}

func (conf *Conf) MapTo(v interface{}) error {
	return conf.iniFile.MapTo(v)
}

func MapTo(conf interface{}, source string) error {
	err := ini.MapTo(conf, source)
	return err
}

func init() {
	ini.DefaultHeader = true
}
