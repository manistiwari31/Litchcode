object Solution {
      import scala.collection.mutable.Stack
    def evalRPN(tokens: Array[String]): Int = {

    val stack = Stack[Int]()
    val op = Set("*", "+", "-", "/")

    for (token <- tokens) {
      if (op.contains(token)) {
        val b = stack.pop()
        val a = stack.pop()

        token match {
          case "*" => stack.push(a * b)
          case "+" => stack.push(a + b)
          case "-" => stack.push(a - b)
          case "/" => stack.push(a / b)  // Note: In Scala, division of integers results in a float
        }
      } else {
        stack.push(token.toInt)
      }
    }

    stack.pop()
  }
}

