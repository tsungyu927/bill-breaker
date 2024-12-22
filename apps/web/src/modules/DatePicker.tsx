import { useEffect, useState } from "react";
import { Calendar } from "@/components/ui/calendar";

interface Props {
  value: Date;
  onChange: (date?: Date) => void;
}

function DatePicker(props: Props) {
  const { value, onChange } = props;
  const [date, setDate] = useState<Date | undefined>(value);

  const handleDateSelect = (selectedDate: Date | undefined) => {
    if (selectedDate) {
      setDate(selectedDate);
    }
  };

  useEffect(() => {
    onChange(date);
  }, [date, onChange]);

  return (
    <div className="flex">
      <Calendar
        mode="single"
        selected={date}
        onSelect={handleDateSelect}
        autoFocus
      />
    </div>
  );
}
export default DatePicker;
