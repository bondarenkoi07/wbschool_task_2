package pattern

import "math"

type Func func(float64) float64

type Equitation interface {
	GetFunc() Func
}

type QuadEquitation struct {
	A, B, C float64
}

type LinearEquitation struct {
	A, B float64
}

type PolynomialEquitation struct {
	args []float64
}

func (p *PolynomialEquitation) SetArgs(args []float64) {
	p.args = args
}

func (p PolynomialEquitation) GetFunc() Func {
	return func(x float64) float64 {
		result := 0.0
		exp := float64(len(p.args) - 1)
		for _, value := range p.args {
			result += math.Pow(x, exp) * value
			exp--
		}
		return result
	}
}

type Strategy interface {
	Calculate() []float64
}

type AnalyticStrategy struct {
	eq QuadEquitation
}

func (a AnalyticStrategy) Calculate() []float64 {
	D := a.eq.B*a.eq.B - 4*a.eq.C*a.eq.A
	x1 := (-a.eq.B + math.Sqrt(D)) / (2 * a.eq.A)
	x2 := (-a.eq.B - math.Sqrt(D)) / (2 * a.eq.A)
	return []float64{x1, x2}
}

type LinearStrategy struct {
	eq LinearEquitation
}

func (a LinearStrategy) Calculate() []float64 {
	return []float64{-a.eq.B / a.eq.A}
}

type CommonStrategy struct {
	eq  PolynomialEquitation
	eps float64
}

func (p CommonStrategy) Calculate() []float64 {
	a := -50.0
	b := 50.0
	f := p.eq.GetFunc()
	for math.Abs(a-b) > p.eps {
		a = b - (b-a)*f(b)/(f(b)-f(a))
		b = a - (a-b)*f(a)/(f(a)-f(b))
	}
	return []float64{b}
}

type StrategyExecutor struct {
	strategy Strategy
}

func (s *StrategyExecutor) SetStrategy(strategy Strategy) {
	s.strategy = strategy
}

func (s StrategyExecutor) Exec() []float64 {
	return s.strategy.Calculate()
}
