export function formatDuration(durationMs: number): string {
  const seconds = Math.floor(durationMs / 1000);
  const minutes = Math.floor(seconds / 60);
  const hours = Math.floor(minutes / 60);
  const days = Math.floor(hours / 24);

  const minutesLeft = minutes % 60;
  const hoursLeft = hours % 24;

  // Check if total hours are less than 24 to decide on including days or not
  if (hours < 24) {
    return `${hoursLeft}h${minutesLeft}m`;
  } else {
    return `${days}d${hoursLeft}h${minutesLeft}m`;
  }
}

export function formatDurationFromStrings(start: string, end: string): string {
  const startDate = new Date(start);
  const endDate = new Date(end);
  const durationMs = endDate.getTime() - startDate.getTime();
  return formatDuration(durationMs);
}
