func bagOfTokensScore(tokens []int, power int) int {
    slices.Sort(tokens)
    var score = 0
    a,b := 0, len(tokens)-1

    for a <=  b {
        if power >= tokens[a]{
            power -= tokens[a]
            a++
            score++
        } else if b >= a+1 && score >0{
            power += tokens[b]
            score--
            b--
        } else {
            break
        }
    }
    return score
}