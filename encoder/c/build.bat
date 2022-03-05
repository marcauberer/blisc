mkdir bin
cd bin

cmake -DCMAKE_C_COMPILER=gcc -DCMAKE_CXX_COMPILER=g++ -DCMAKE_C_FLAGS="-O2" -GNinja ..
cmake --build .

cd ..