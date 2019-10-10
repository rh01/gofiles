/*
 * Copyright 2018 @rh01 https://github.com/rh01
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *    http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */
package lcode.ex03;

import java.util.Stack;

/**
 * Date 2019/10/10
 *
 * @author: Shen Hengheng
 * <p>
 * <p>
 * <p>
 * <p>
 * <p>
 * <p>
 * Solution:
 * </p>
 **/
public class Solution {
    public  int evalRPN(String[] tokens) {
        Stack<Integer> stack = new Stack<>();

        if (tokens == null || tokens.length == 0) {
            return 0;
        }

        int r, l, res = 0;
        for (String s : tokens) {
            switch (s) {
                case "/":
                    r = stack.pop();
                    l = stack.pop();
                    res = l / r;
                    stack.push(res);
                    break;

                case "*":
                    r = stack.pop();
                    l = stack.pop();
                    res = l * r;
                    stack.push(res);
                    break;

                case "-":
                    r = stack.pop();
                    l = stack.pop();
                    res = l - r;
                    stack.push(res);
                    break;

                case "+":
                    r = stack.pop();
                    l = stack.pop();
                    res = l + r;
                    stack.push(res);
                    break;
                default:
                    stack.push(Integer.parseInt(s));
            }
        }

        return stack.pop();

    }
}