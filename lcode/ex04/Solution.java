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
package lcode.ex04;

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
// Definition for a point.

import java.util.HashMap;
import java.util.Map;

class Solution {
    public int maxPoints(int[][] g) {
        int n = g.length;
        if (n <= 2) return n;

        int max = 0;

        // slope (dy / dx) --> count
        Map<Integer, Map<Integer, Integer>> map = new HashMap();
        for (int i = 0; i < g.length; ++i) {
            map.clear();
            int same = 0;
            int line_max_c = 0;
            for (int j = 0; j < g.length; ++j) {
                if (is_same(g[i], g[j])) {
                    ++same;
                } else {
                    int dy = g[j][1] - g[i][1];
                    int dx = g[j][0] - g[i][0];
                    int line_c = put(map, dy, dx);
                    line_max_c = Math.max(line_c, line_max_c);
                }
            }

            max = Math.max(max, same + line_max_c);
        }

        return max;
    }

    boolean is_same(int[] a, int[] b) {
        return a[0] == b[0] && a[1] == b[1];
    }

    // slope (dy / dx) --> count
    int put(Map<Integer, Map<Integer, Integer>> map, int dy, int dx) {
        if (dx == 0) dy = 1;
        else if (dy == 0) dx = 1;
        else {
            int sign = ((dy > 0 && dx > 0) || (dy < 0 && dx < 0)) ? 1 : -1;
            dy = Math.abs(dy);
            dx = Math.abs(dx);
            int c = gcd(dy, dx);
            dy /= c;
            dx /= c;
            if (sign < 0) dy = -dy;
        }

        Map<Integer, Integer> dx_c = map.get(dy);
        if (dx_c == null) {
            dx_c = new HashMap();
            dx_c.put(dx, 1);
            map.put(dy, dx_c);
            return 1;
        } else {
            dx_c.put(dx, 1 + dx_c.getOrDefault(dx, 0));
            return dx_c.get(dx);
        }
    }

    // assume: a >= b >= 0
    int gcd(int a, int b) {
        if (a < b) return gcd(b, a);
        if (b == 0) return a;
        return gcd(b, a % b);
    }
}
