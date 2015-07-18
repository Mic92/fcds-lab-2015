set datafile separator ","
set terminal png size 1024,768 enhanced font "sans,15"

set xlabel "Number of Cores"
set ylabel "Time[s]"

set output "../pages/presentation/threesat-final.png"
set title "3Sat [File: large.in]"

plot "threesat-final.data" using 1:($3/1000) title "Go: Total time ≈ Computation Time" with points ps 3, \
  "threesat-final.data" using 1:($3/1000):(sprintf("%.2f%%\n%.2fs", 30532.125/$2 * 100, $2/1000)) with labels offset character 0, character 2 notitle, \
