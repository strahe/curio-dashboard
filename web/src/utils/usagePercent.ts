export function usagePercent(available: number, capacity: number, decimalPlaces: number = 2) {
  const percent = 100 - (available / capacity) * 100;
  return parseFloat(percent.toFixed(decimalPlaces));
}
