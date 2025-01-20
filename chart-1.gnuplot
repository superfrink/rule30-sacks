#set term png
#set term png size 15000,15000
set output "output.png"
#set grid

#plot [-15:15] [-15:15] "coordinates.csv" using 1:2 "%lf,%lf"
#plot [-50:50] [-50:50] "coordinates.csv" using 1:2 "%lf,%lf" with dots
#plot [-50:50] [-50:50] "coordinates.csv" using 1:2 "%lf,%lf" pt 7 ps 4

#set term png size 15000,15000
#plot [-1000:1000] [-1000:1000] "coordinates.csv" using 1:2 "%lf,%lf" pt 7 ps 1

set term png size 10000,10000
plot [-500:500] [-500:500] "coordinates.csv" using 1:2 "%lf,%lf" pt 7 ps 2


