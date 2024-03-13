#include <stdint.h>
#include <stdio.h>
extern int64_t oSumTo100(int64_t i);
extern int64_t sumTo100(int64_t i);
struct tuple2float_int{
double x;
long long y;
};
extern struct tuple2float_int reReTuple(long long, double);
extern int testReTuple();

struct tuple2float_int foo(){
struct tuple2float_int xx={
.x=4,
.y=6};
return xx;
}

int main(int argc, const char* argv[]){

testReTuple();
return 0;
}