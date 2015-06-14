set datafile separator ","
set terminal png size 1024,768 enhanced font "sans"

set xlabel "Number of Cores"
set ylabel "Time[s]"

set output '../presentation/threesat.png'
set title "Gnuplot"
ax=1
ay=63.428
plot 'threesat.data' using 1:($2/1000) title "Go" pt 7 ps 2, \
       "<echo 1 1" using (ax):(ay)  title "C Reference Implementation" pt 7 ps 2
