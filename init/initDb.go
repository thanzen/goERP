//数据库部分数据初始化
package init

import (
	"encoding/xml"
	"io/ioutil"
	"os"
	. "pms/models/base"
	"runtime"
)

//初始化数据库
func InitDb() {
	systemType := runtime.GOOS
	split := "/"
	switch systemType {
	case "windows":
		split = "\\"
	case "linux":
		split = "/"
	}
	if xmDir, err := os.Getwd(); err == nil {
		if _, err := GetUser(1); err != nil {

			xmDir += split + "init_xml" + split
			initUser(xmDir + "Users.xml")
			if user, err := GetUser(1); err == nil {
				initGroup(xmDir+"Groups.xml", user)
				initCountry(xmDir+"Countries.xml", user)
				initProvince(xmDir+"Provinces.xml", user)
				initCity(xmDir+"Cities.xml", user)
				initDistrict(xmDir+"Districts.xml", user)
			}
		}
	}

}

func initUser(filename string) {
	if file, err := os.Open(filename); err == nil {
		defer file.Close()
		if data, err := ioutil.ReadAll(file); err == nil {
			var initUsers InitUsers
			if xml.Unmarshal(data, &initUsers) == nil {
				for _, k := range initUsers.Users {
					//admin系统管理员
					if k.Name == "admin" {
						k.Id = 1
					}
					AddUser(k, k)
				}
			}
		}
	}

}
func initGroup(filename string, user User) {
	if file, err := os.Open(filename); err == nil {
		defer file.Close()
		if data, err := ioutil.ReadAll(file); err == nil {
			var initGroups InitGroups
			if xml.Unmarshal(data, &initGroups) == nil {
				for _, k := range initGroups.Groups {
					AddGroup(k, user)
				}
			}
		}
	}

}

func initCountry(filename string, user User) {
	if file, err := os.Open(filename); err == nil {
		defer file.Close()
		if data, err := ioutil.ReadAll(file); err == nil {
			var initCountries InitCountries
			if xml.Unmarshal(data, &initCountries) == nil {
				for _, k := range initCountries.Countries {
					AddCountry(k, user)
				}
			}
		}
	}
}
func initProvince(filename string, user User) {
	if file, err := os.Open(filename); err == nil {
		defer file.Close()
		if data, err := ioutil.ReadAll(file); err == nil {
			var initProvinces InitProvinces
			if xml.Unmarshal(data, &initProvinces) == nil {
				for _, k := range initProvinces.Provinces {
					var province Province
					pid := int64(k.PID)
					if country, err := GetCountry(pid); err == nil {
						province.Country = &country
						province.Name = k.Name
						AddProvince(province, user)
					}
				}
			}
		}
	}
}
func initCity(filename string, user User) {
	if file, err := os.Open(filename); err == nil {
		defer file.Close()
		if data, err := ioutil.ReadAll(file); err == nil {
			var initCities InitCities
			if xml.Unmarshal(data, &initCities) == nil {
				for _, k := range initCities.Cities {
					var city City
					pid := int64(k.PID)
					if province, err := GetProvince(pid); err == nil {
						city.Province = &province
						city.Name = k.Name
						AddCity(city, user)
					}
				}
			}
		}
	}
}
func initDistrict(filename string, user User) {
	if file, err := os.Open(filename); err == nil {
		defer file.Close()
		if data, err := ioutil.ReadAll(file); err == nil {
			var initDistricts InitDistricts
			if xml.Unmarshal(data, &initDistricts) == nil {
				for _, k := range initDistricts.Districts {
					var district District
					pid := int64(k.PID)
					if city, err := GetCity(pid); err == nil {
						district.City = &city
						district.Name = k.Name
						AddDistrict(district, user)
					}
				}
			}
		}
	}
}
