package backoff

// Null creates a new NullPolicy object
func Null() Policy {
	return NewNull()
}

// Constant creates a new ConstantPolicy object
func Constant(options ...Option) Policy {
	return NewConstantPolicy(options...)
}

// Constant creates a new ExponentialPolicy object
func Exponential(options ...Option) Policy {
	return NewExponentialPolicy(options...)
}

// Continue is a convenience function to check when we can fire
// the next invocation of the desired backoff code
//
// for backoff.Continue(c) {
//  ... your code ...
// }
func Continue(c Controller) bool {
	select {
	case <-c.Done():
		return false
	case <-c.Next():
		return true
	}
}
