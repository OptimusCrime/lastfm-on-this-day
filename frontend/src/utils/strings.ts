export const addLeadingZero = (value: number): string => (value < 10 ? `0${value}` : value.toString());
