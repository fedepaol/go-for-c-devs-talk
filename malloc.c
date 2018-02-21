int* bird = malloc(sizeof(int)) // lives in the heap, need to be freed

int* f() {
    int b = 5;
    return &b; // BAD!
}
