export const combineDateTime = (date: Date, time: Date): Date => {
  const newDate = new Date(date);

  newDate.setHours(time.getHours());
  newDate.setMinutes(time.getMinutes());

  return newDate;
};
