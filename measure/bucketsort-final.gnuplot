set datafile separator ","
set terminal png size 1024,768 enhanced font "sans,15"

set xlabel "Number of Cores"
set ylabel "Time[ms]/Speedup[%]"

set output "../pages/presentation/bucketsort-final.png"
set title "Bucketsort [File: large1.in]"

plot "bucketsort-final.data" using 1:2 title "Go: Computation time" with points ps 3, \
  "bucketsort-final.data" using 1:2:(sprintf("%.2f%%\n%.2fms", 536.5/$2 * 100, $2)) with labels offset character 0, character 2 notitle,  \
  "bucketsort-final.data" using 1:3 title "Go: Total Time" with points ps 3, \
  "bucketsort-final.data" using 1:3:(sprintf("%.2f%%\n%.2fms", 1155.25/$3 * 100, $3)) with labels offset character 0, character 2 notitle
