Arrays.stream(myArray)
	      .filter(s -> s.length() > 4)
	      .map(s -> s.toUpperCase())
	      .toArray(String[]::new);
