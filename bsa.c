#include <stdio.h>

int in_num, out_num;

int calc_fa_num(int* state) {
    int fa_num = 0;
    for (int i = out_num - 1; i >= 0; i--) {
         int quotient = state[i] / 3;
         int remainder = state[i] % 3;
         fa_num += quotient;
         state[i] = quotient + remainder;
         if (i != out_num -1) {
             state[i+1] += quotient;
         }
    }   
    return fa_num;
}

int main() {
    scanf("%d%d", &in_num, &out_num);
    int num_state[out_num];
    for (int i = 0; i < out_num; i++) {
        num_state[i] = 0;
    }
    num_state[0] = in_num;

    int step_num = 0;
    int fa_num = 0;
    while (num_state[out_num - 1] == 0) {
        step_num++;
        fa_num += calc_fa_num(num_state);
    }
    printf("step_num: %d, fa_num: %d\n", step_num, fa_num);
    return 0;
}
