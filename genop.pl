#!/usr/bin/env perl

use strict;
use warnings;
use v5.28;

my %dyads = (
    Equal =>  {
        B_B => ["w == x", "B"],
        B_I => ["B2I(w) == x", "B"],
        B_F => ["B2F(w) == x", "B"],
        I_B => ["w == B2I(x)", "B"],
        I_I => ["w == x", "B"],
        I_F => ["F(w) == x", "B"],
        F_B => ["w == B2F(x)", "B"],
        F_I => ["w == F(x)", "B"],
        F_F => ["w == x", "B"],
        S_S => ["w == x", "B"],
    },
    NotEqual =>  {
        B_B => ["w != x", "B"],
        B_I => ["B2I(w) != x", "B"],
        B_F => ["B2F(w) != x", "B"],
        I_B => ["w != B2I(x)", "B"],
        I_I => ["w != x", "B"],
        I_F => ["F(w) != x", "B"],
        F_B => ["w != B2F(x)", "B"],
        F_I => ["w != F(x)", "B"],
        F_F => ["w != x", "B"],
        S_S => ["w != x", "B"],
    },
    Lesser =>  {
        B_B => ["B2I(w) < B2I(x)", "B"],
        B_I => ["B2I(w) < x", "B"],
        B_F => ["B2F(w) < x", "B"],
        I_B => ["w < B2I(x)", "B"],
        I_I => ["w < x", "B"],
        I_F => ["F(w) < x", "B"],
        F_B => ["w < B2F(x)", "B"],
        F_I => ["w < F(x)", "B"],
        F_F => ["w < x", "B"],
        S_S => ["w < x", "B"],
    },
    LesserEq =>  {
        B_B => ["B2I(w) <= B2I(x)", "B"],
        B_I => ["B2I(w) <= x", "B"],
        B_F => ["B2F(w) <= x", "B"],
        I_B => ["w <= B2I(x)", "B"],
        I_I => ["w <= x", "B"],
        I_F => ["F(w) <= x", "B"],
        F_B => ["w <= B2F(x)", "B"],
        F_I => ["w <= F(x)", "B"],
        F_F => ["w <= x", "B"],
        S_S => ["w <= x", "B"],
    },
    Greater =>  {
        B_B => ["B2I(w) > B2I(x)", "B"],
        B_I => ["B2I(w) > x", "B"],
        B_F => ["B2F(w) > x", "B"],
        I_B => ["w > B2I(x)", "B"],
        I_I => ["w > x", "B"],
        I_F => ["F(w) > x", "B"],
        F_B => ["w > B2F(x)", "B"],
        F_I => ["w > F(x)", "B"],
        F_F => ["w > x", "B"],
        S_S => ["w > x", "B"],
    },
    GreaterEq =>  {
        B_B => ["B2I(w) >= B2I(x)", "B"],
        B_I => ["B2I(w) >= x", "B"],
        B_F => ["B2F(w) >= x", "B"],
        I_B => ["w >= B2I(x)", "B"],
        I_I => ["w >= x", "B"],
        I_F => ["F(w) >= x", "B"],
        F_B => ["w >= B2F(x)", "B"],
        F_I => ["w >= F(x)", "B"],
        F_F => ["w >= x", "B"],
        S_S => ["w >= x", "B"],
    },
    Add =>  {
        B_B => ["B2I(w) + B2I(x)", "I"],
        B_I => ["B2I(w) + x", "I"],
        B_F => ["B2F(w) + x", "F"],
        I_B => ["w + B2I(x)", "I"],
        I_I => ["w + x", "I"],
        I_F => ["F(w) + x", "F"],
        F_B => ["w + B2F(x)", "F"],
        F_I => ["w + F(x)", "F"],
        F_F => ["w + x", "F"],
    },
    Subtract =>  {
        B_B => ["B2I(w) - B2I(x)", "I"],
        B_I => ["B2I(w) - x", "I"],
        B_F => ["B2F(w) - x", "F"],
        I_B => ["w - B2I(x)", "I"],
        I_I => ["w - x", "I"],
        I_F => ["F(w) - x", "F"],
        F_B => ["w - B2F(x)", "F"],
        F_I => ["w - F(x)", "F"],
        F_F => ["w - x", "F"],
    },
    Multiply =>  {
        B_B => ["w && x", "B"],
        B_I => ["B2I(w) * x", "I"],
        B_F => ["B2F(w) * x", "F"],
        I_B => ["w * B2I(x)", "I"],
        I_I => ["w * x", "I"],
        I_F => ["F(w) * x", "F"],
        F_B => ["w * B2F(x)", "F"],
        F_I => ["w * F(x)", "F"],
        F_F => ["w * x", "F"],
    },
    Divide =>  {
        B_B => ["divide(B2F(w), B2F(x))", "F"],
        B_I => ["divide(B2F(w), F(x))", "F"],
        B_F => ["divide(B2F(w), x)", "F"],
        I_B => ["divide(F(w), B2F(x))", "F"],
        I_I => ["divide(F(w), F(x))", "F"],
        I_F => ["divide(F(w), x)", "F"],
        F_B => ["divide(w, B2F(x))", "F"],
        F_I => ["divide(w, F(x))", "F"],
        F_F => ["divide(w, x)", "F"],
    },
    Minimum =>  {
        B_B => ["w && x", "B"],
        B_I => ["minI(B2I(w), x)", "I"],
        B_F => ["F(math.Min(float64(B2F(w)), float64(x)))", "F"],
        I_B => ["minI(w, B2I(x))", "I"],
        I_I => ["minI(w, x)", "I"],
        I_F => ["F(math.Min(float64(w), float64(x)))", "F"],
        F_B => ["F(math.Min(float64(w), float64(B2F(x))))", "F"],
        F_I => ["F(math.Min(float64(w), float64(x)))", "F"],
        F_F => ["F(math.Min(float64(w), float64(x)))", "F"],
        S_S => ["minS(w, x)", "S"],
    },
    Maximum =>  {
        B_B => ["w || x", "B"],
        B_I => ["maxI(B2I(w), x)", "I"],
        B_F => ["F(math.Max(float64(B2F(w)), float64(x)))", "F"],
        I_B => ["maxI(w, B2I(x))", "I"],
        I_I => ["maxI(w, x)", "I"],
        I_F => ["F(math.Max(float64(w), float64(x)))", "F"],
        F_B => ["F(math.Max(float64(w), float64(B2F(x))))", "F"],
        F_I => ["F(math.Max(float64(w), float64(x)))", "F"],
        F_F => ["F(math.Max(float64(w), float64(x)))", "F"],
        S_S => ["maxS(w, x)", "S"],
    },
    Or =>  {
        B_B => ["w || x", "B"],
        B_I => ["1-((1-B2I(w)) * (1-x))", "I"],
        B_F => ["1-((1-B2F(w)) * (1-x))", "F"],
        I_B => ["1-((1-w) * (1-B2I(x)))", "I"],
        I_I => ["1-((1-w) * (1-x))", "I"],
        I_F => ["1-((1-F(w)) * (1-x))", "F"],
        F_B => ["1-((1-w) * (1-B2F(x)))", "F"],
        F_I => ["1-((1-w) * F(1-x))", "F"],
        F_F => ["1-((1-w) * (1-x))", "F"],
    },
    And =>  {
        B_B => ["w && x", "B"],
        B_I => ["B2I(w) * x", "I"],
        B_F => ["B2F(w) * x", "F"],
        I_B => ["w * B2I(x)", "I"],
        I_I => ["w * x", "I"],
        I_F => ["F(w) * x", "F"],
        F_B => ["w * B2F(x)", "F"],
        F_I => ["w * F(x)", "F"],
        F_F => ["w * x", "F"],
    },
);

print <<EOS;
package main

import "math"

EOS

genOp("Equal", "=");
genOp("NotEqual", "≠");
genOp("Lesser", "<");
genOp("LesserEq", "≤");
genOp("Greater", ">");
genOp("GreaterEq", "≥");
genOp("Add", "+");
genOp("Subtract", "-");
genOp("Multiply", "×");
genOp("Divide", "÷");
genOp("Minimum", "⌊");
genOp("Maximum", "⌈");
genOp("And", "∧"); # identical to Multiply
genOp("Or", "∨"); # Multiply under Not

sub genOp {
    my ($name, $op) = @_;
    my $cases = $dyads{$name};
    my %types = map { $_=~/(\w)_/; $1 => 1 } keys $cases->%*;
    my $s = "";
    open my $out, '>', \$s;
    print $out <<EOS;
func ${name}(w, x O) O {
	switch w := w.(type) {
EOS
    for my $t (sort keys %types) {
        print $out <<EOS;
	case $t:
		return ${name}${t}O(w, x)
EOS
    }
    for my $t (sort keys %types) {
        print $out <<EOS;
	case A$t:
		return ${name}A${t}O(w, x)
EOS
    }
    print $out <<EOS;
	case AO:
		r := make(AO, len(w))
		for i := range r {
			v := ${name}(w[i], x)
			e, ok := v.(E)
			if ok {
				return e
			}
			r[i] = v
		}
		return r
	case E:
		return w
	default:
		return badtype("${op}")
	}
}\n
EOS
    print $s;
    for my $t (sort keys %types) {
        genLeftExpanded($name, $op, $cases, $t);
    }
    for my $t (sort keys %types) {
        genLeftArrayExpanded($name, $op, $cases, $t);
    }
}

sub genLeftExpanded {
    my ($name, $op, $cases, $t) = @_;
    my %types = map { /_(\w)/; $1 => $cases->{"${t}_$1"}} grep { /${t}_(\w)/ } keys $cases->%*;
    my $s = "";
    open my $out, '>', \$s;
    print $out <<EOS;
func ${name}${t}O(w $t, x O) O {
	switch x := x.(type) {
EOS
    for my $tt (sort keys %types) {
        my $expr = $cases->{"${t}_$tt"}->[0];
        print $out <<EOS;
	case $tt:
		return $expr
EOS
    }
    for my $tt (sort keys %types) {
        my $expr = $cases->{"${t}_$tt"}->[0];
        my $type = $cases->{"${t}_$tt"}->[1];
        my $iexpr = subst($expr, "w", "x[i]");
        print $out <<EOS;
	case A$tt:
		r := make(A$type, len(x))
		for i := range r {
			r[i] = $iexpr
		}
		return r
EOS
    }
    print $out <<EOS if $t !~ /^A/;
	case AO:
		r := make([]O, len(x))
		for i := range r {
			v := ${name}${t}O(w, x[i])
			e, ok := v.(E)
			if ok {
				return e
			}
			r[i] = v
		}
		return r
	case E:
		return w
	default:
		return badtype("${op}")
	}
}\n
EOS
    print $s;
}

sub genLeftArrayExpanded {
    my ($name, $op, $cases, $t) = @_;
    my %types = map { /_(\w)/; $1 => $cases->{"${t}_$1"}} grep { /${t}_(\w)/ } keys $cases->%*;
    my $s = "";
    open my $out, '>', \$s;
    print $out <<EOS;
func ${name}A${t}O(w A$t, x O) O {
	switch x := x.(type) {
EOS
    for my $tt (sort keys %types) {
        my $expr = $cases->{"${t}_$tt"}->[0];
        my $type = $cases->{"${t}_$tt"}->[1];
        my $iexpr = subst($expr, "w[i]", "x");
        print $out <<EOS;
	case $tt:
		r := make(A$type, len(w))
		for i := range r {
			r[i] = $iexpr
		}
		return r
EOS
    }
    for my $tt (sort keys %types) {
        my $expr = $cases->{"${t}_$tt"}->[0];
        my $type = $cases->{"${t}_$tt"}->[1];
        my $iexpr = subst($expr, "w[i]", "x[i]");
        print $out <<EOS;
	case A$tt:
		r := make(A$type, len(x))
		for i := range r {
			r[i] = $iexpr
		}
		return r
EOS
    }
    print $out <<EOS if $t !~ /^A/;
	case AO:
		r := make(AO, len(x))
		for i := range r {
			v := ${name}A${t}O(w, x[i])
			e, ok := v.(E)
			if ok {
				return e
			}
			r[i] = v
		}
		return r
	case E:
		return w
	default:
		return badtype("${op}")
	}
}\n
EOS
    print $s;
}

sub subst {
    my ($expr, $w, $x) = @_;
    $expr =~ s/\bw\b/$w/g;
    $expr =~ s/\bx\b/$x/g;
    return $expr
}
