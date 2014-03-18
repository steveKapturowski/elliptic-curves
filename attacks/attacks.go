/*
This library is intended ONLY for the safety verification of cryptographic
schemes based on the discrete logarithm problem, and is not in any way
intended to aid attackers in breaking cryptographic schemes that are used
in practice. Furthermore, we take the the view that any scheme utilizing 
a set of parameters that CAN be broken by the below algorithms has no 
business being used for any practical security purposes.
*/

package attacks

import(
	"math/big"
	"../../HyperellipticFuntime"
)

// TODO: generalize attacks where possible to work on any group

// Square root algorithms
func PollardRho(curve *hyperelliptic.HyperellipticCurve, 
				basePoint, exponential *hyperelliptic.Point) *big.Int {
	return new(big.Int)
}

func BabyStepGiantStep(curve hyperelliptic.HyperellipticCurve,
					   basePoint, exponential *hyperelliptic.Point) *big.Int {
	return new(big.Int)
}

/* Index calculus algorithms  
time complexity: exp( (sqrt(2)+O(1)) * sqrt(ln(p)lnln(p)) )
*/

