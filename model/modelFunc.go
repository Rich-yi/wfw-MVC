package model

func GetAreas()([]Area,error)  {
	var areas []Area
if err:=	GlobalDB.Find(&areas).Error;err!=nil{
	return nil,err
}
return areas,nil

}