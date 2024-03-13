%IntFloat = type { i64, double }

@.str.1 = global [21 x i8] c"result is %lf, %lld\0A\00"

define i64 @max(i64 %0, i64 %1) {
entry:
	%2 = icmp sgt i64 %0, %1
	br i1 %2, label %3, label %4

3:
	br label %5

4:
	br label %5

5:
	%6 = phi i64 [ %0, %3 ], [ %1, %4 ]
	ret i64 %6
}

define i64 @max3(i64 %0, i64 %1, i64 %2) {
entry:
	%3 = call i64 @max(i64 %0, i64 %1)
	%4 = call i64 @max(i64 %3, i64 %2)
	ret i64 %4
}

define i64 @sumTo100(i64 %0) {
entry:
	%1 = icmp sgt i64 %0, 100
	br i1 %1, label %2, label %3

2:
	br label %7

3:
	%4 = add i64 %0, 1
	%5 = call i64 @sumTo100(i64 %4)
	%6 = add i64 %0, %5
	br label %7

7:
	%8 = phi i64 [ 0, %2 ], [ %6, %3 ]
	ret i64 %8
}

define i64 @oSumTo100(i64 %0, i64 %1) {
entry:
	%2 = icmp sgt i64 %0, 100
	br i1 %2, label %3, label %4

3:
	br label %8

4:
	%5 = add i64 %0, 1
	%6 = add i64 %1, %0
	%7 = call i64 @oSumTo100(i64 %5, i64 %6)
	br label %8

8:
	%9 = phi i64 [ %1, %3 ], [ %7, %4 ]
	ret i64 %9
}

declare i64 @printf(i8* %0, double %1, i64 %2)

define { double, i64 } @reTuple(i64 %0, double %1) {
entry:
	%2 = insertvalue { double, i64 } undef, double %1, 0
	%3 = insertvalue { double, i64 } %2, i64 %0, 1
	ret { double, i64 } %3
}

define { i64, double } @re1() {
entry:
	%0 = insertvalue { i64, double } undef, i64 1, 0
	%1 = insertvalue { i64, double } %0, double 0x3FF3333333333333, 1
	ret { i64, double } %1
}

define { double, i64 } @re2() {
entry:
	%0 = call { i64, double } @re1()
	%1 = extractvalue { i64, double } %0, 0
	%2 = extractvalue { i64, double } %0, 1
	%3 = call { double, i64 } @reTuple(i64 %1, double %2)
	ret { double, i64 } %3
}

define i64 @re3(i64 %0, { i64, i64 } %1) {
entry:
	%2 = extractvalue { i64, i64 } %1, 1
	%3 = icmp sgt i64 %2, 0
	br i1 %3, label %4, label %7

4:
	%5 = mul i64 3, %0
	%6 = mul i64 %5, %0
	br label %8

7:
	br label %8

8:
	%9 = phi i64 [ %6, %4 ], [ 0, %7 ]
	ret i64 %9
}

define i64 @testReTuple() {
entry:
	%0 = getelementptr [21 x i8], [21 x i8]* @.str.1, i64 0, i64 0
	%1 = call { double, i64 } @reTuple(i64 2, double 0x400B333333333333)
	%2 = extractvalue { double, i64 } %1, 0
	%3 = call { double, i64 } @reTuple(i64 2, double 0x400B333333333333)
	%4 = extractvalue { double, i64 } %3, 1
	%5 = call i64 @printf(i8* %0, double %2, i64 %4)
	ret i64 %5
}
