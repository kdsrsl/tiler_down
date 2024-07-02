# geojson
2024年06月25日11:00:07



1. https://geojson.io/#map=2/0/20 可以在线框选
1. https://datav.aliyun.com/portal/school/atlas/area_selector  最小单位是县
2. https://www.poi86.com/



```
SENRSL:tiler_down senrsl$ export https_proxy=http://192.168.0.148:7890 http_proxy=http://192.168.0.148:7890 all_proxy=socks5://192.168.0.148:7891
SENRSL:tiler_down senrsl$ 
SENRSL:tiler_down senrsl$ go run main.go map.go task.go tile.go utils.go
2024-06-25 11:16:00.371 [INFO] zoom: 16, tiles: 33 

2024-06-25 11:16:00.372 [INFO] zoom: 17, tiles: 63 

2024-06-25 11:16:00.372 [INFO] zoom: 18, tiles: 127 

2024-06-25 11:16:00.372 [INFO] zoom: 19, tiles: 251 

2024-06-25 11:16:00.372 [INFO] zoom: 20, tiles: 502 

2024-06-25 11:16:00.372 [INFO] zoom: 21, tiles: 1004 

2024-06-25 11:16:00.372 [INFO] zoom: 22, tiles: 2008 
```


使用

```
SENRSL:tiler_down senrsl$ go run api4.go 
```


## 资料

1. http://webgis.cn/geog585/ch05_buildingtiled-maps-with-foss/sec01_overview/section.html
2. java下载 + vue展示 https://blog.csdn.net/hualinger/article/details/134787040
3. https://github.com/nextprops/note/blob/master/map/地图相关知识整理(瓦片地图).md
4. https://mapshaper.org/ 这个是shp格式的在线预览吧


