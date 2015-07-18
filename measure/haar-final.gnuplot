set datafile separator ","
set terminal png size 1024,768 enhanced font "sans,15"

set xlabel "Number of Cores"
set ylabel "Time[ms]/Speedup[%]"

set output "../pages/presentation/haar-final.png"
set title "Haar Wavelets [File: large.in]"
plot "haar-final.data" using 1:2 title "Go: Computation time" with points ps 3, \
  "haar-final.data" using 1:2:(sprintf("%.2f%%\n%.2fms", 2436.75/$2 * 100, $2)) with labels offset character 0, character 2 notitle, \
  "haar-final.data" using 1:3 title "Go: Total Time" with points ps 3, \
  "haar-final.data" using 1:3:(sprintf("%.2f%%\n%.2fms", 3844.875/$3 * 100, $3)) with labels offset character 0, character 2 notitle
