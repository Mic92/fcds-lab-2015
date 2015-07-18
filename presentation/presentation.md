<!--
Multi-Monitor-Shortcuts:
Ctrl-O: Move Window to next screen
Mod4 + Control + j/k: Focus next/previous screen
reveal.js-Shortcuts:
o: Öffne Übersicht
s: Öffne Vortragsmonitor
-->

<!-- .slide: data-state="intro" -->
# FCDS Lab 2015

Jörg Thalheim

<joerg@higgsboson.tk>



## Algorithms
<p>
<input type="checkbox" checked> bucketsort<br>
<input type="checkbox"> friendly<br>
<input type="checkbox" checked> haar<br>
<input type="checkbox"> knapsack<br>
<input type="checkbox" checked> threesat
</p>




## Bucketsort: Overview
<img src="bucketsort-overview.png" alt="bucketsort" height="500">



## Haar Wavelets: Overview 1
<img src="haar-wavelets-overview1.png" alt="haar wavelets" height="500">



## Haar Wavelets: Overview 2
<img src="haar-wavelets-overview2.png" alt="haar wavelets" height="500">



## Threesat: Overview
<img src="threesat-overview1.png" alt="threesat" height="500">



## Measurement
- for `1` to `8` cpu cores
  - take `10` measurements
    - take computation time (only core algorithm)
    - take total run time (computation time + time spent on I/O)
  - exclude the slowest and fastest measurement
  - average over the remaining `8` measurements
- after each measurement call `sync` command, to flush remaining I/O buffers of the
  operating system



## Bucketsort: Performance
<img src="bucketsort-final.png" alt="bucketsort" height="500">



## Threesat: Performance
<img src="threesat-final.png" alt="threesat" height="500">



## Haar Wavelets: Performance
<img src="haar-final.png" alt="threesat" height="500">



## Summary
