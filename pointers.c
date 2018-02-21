int i = 5;

void changeValue(int* i)
{
    *i = 12;
}

int array[3] = {0, 1, 2};

printf("%d", *(array + 2 * sizeof(int)); // 2
