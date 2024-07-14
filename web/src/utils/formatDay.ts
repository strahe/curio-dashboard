export function formatDay(date: Date): string {
  const month = (date.getMonth() + 1).toString().padStart(2, '0'); // Pad with leading zero if needed
  const day = date.getDate().toString().padStart(2, '0'); // Pad with leading zero if needed
  return `${month}-${day}`;
}
