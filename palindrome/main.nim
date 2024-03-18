proc isPalindrome(x: int): bool =
    var x:int
    # Negative numbers are not palindromes
    if x < 0:
        return false

    var original = x
    var reversed = 0

    # Reverse the digits of the number
    while x > 0:
        let digit = x mod 10
        reversed = reversed * 10 + digit
        x = x div 10

    return original == reversed

# Example usage
echo isPalindrome(121)
