import dayjs from 'dayjs';
import relativeTime from 'dayjs/plugin/relativeTime';

dayjs.extend(relativeTime);

export const fromDate = (date: Parameters<typeof dayjs>[0], relative = new Date()) => {
  return dayjs(date).from(relative);
} 