version: "0.1"
database:
  # consult[https://gorm.io/docs/connecting_to_the_database.html]"
  dsn : "root:123456@(127.0.0.1:3306)/activities?charset=utf8mb4&parseTime=true&loc=Local"
  # input mysql or postgres or sqlite or sqlserver. consult[https://gorm.io/docs/connecting_to_the_database.html]
  db  : "mysql"
  # enter the required data table or leave it blank.You can input : orders,users,goods
  tables  : ""
  # specify a directory for output
  outPath :  "./v2/dao"
  # query code file name, default: gen.go
  outFile :  ""
  # generate unit test for query code
  withUnitTest  : false
  # generated model code's package name
  modelPkgName  : ""
  # generate with pointer when field is nullable
  fieldNullable : false
  # generate field with gorm index tag
  fieldWithIndexTag : false
  # generate field with gorm column type tag
  fieldWithTypeTag  : false