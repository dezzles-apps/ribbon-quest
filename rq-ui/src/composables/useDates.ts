import dayjs from 'dayjs';
import localizedFormat from 'dayjs/plugin/localizedFormat'
dayjs.extend(localizedFormat);

export function useDates() {
  function toLocal(input: string) : string {
    if (!input) {
      return '';
    }
    const date = dayjs(input);
    return date.format("DD-MM-YYYY HH:mm");;
  }

  return { toLocal }
}
