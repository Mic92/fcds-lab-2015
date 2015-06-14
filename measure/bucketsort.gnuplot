set datafile separator ","
set terminal png size 1024,768 enhanced font "sans"

set xlabel "Number of Cores"
set ylabel "Time[s]"

set output '../presentation/bucketsort1.png'
set title "Bucketsort"
ax=1
ay=138156/1000
plot 'bucketsort.data' using 1:($2/1000) title "Go" pt 7 ps 2, \
       "<echo 1 1" using (ax):(ay)  title "C Reference Implementation" pt 7 ps 2

set ylabel "Time[ms]"
set output '../presentation/bucketsort2.png'
plot 'bucketsort.data' using 1:2 title "Go" pt 7 ps 2
