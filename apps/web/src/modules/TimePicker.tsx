import { Button } from "@/components/ui/button";
import { ScrollArea, ScrollBar } from "@/components/ui/scroll-area";
import { useEffect, useRef, useState } from "react";

interface Props {
  value: Date;
  onChange: (date?: Date) => void;
}

function TimePicker(props: Props) {
  const { value, onChange } = props;
  const [time, setTime] = useState<Date | undefined>(value);

  // Refs for the currently selected hour and minute
  const hourRef = useRef<HTMLButtonElement | null>(null);
  const minuteRef = useRef<HTMLButtonElement | null>(null);

  const hours = Array.from({ length: 24 }, (_, i) => i);

  const handleTimeChange = (type: "hour" | "minute", value: string) => {
    if (time) {
      const newDate = new Date(time);
      if (type === "hour") {
        newDate.setHours(parseInt(value));
      } else if (type === "minute") {
        newDate.setMinutes(parseInt(value));
      }
      setTime(newDate);
    }
  };

  useEffect(() => {
    const scrollToRef = (ref: React.RefObject<HTMLButtonElement>) => {
      if (ref.current) {
        try {
          ref.current.scrollIntoView({
            behavior: "smooth",
            block: "center",
          });
        } catch {
          // Fallback for iOS or unsupported environments
          ref.current.scrollIntoView();
        }
      }
    };

    // Scroll to the selected hour
    if (hourRef.current) {
      scrollToRef(hourRef);
    }

    // Scroll to the selected minute
    if (minuteRef.current) {
      scrollToRef(minuteRef);
    }
  }, []);

  useEffect(() => {
    onChange(time);
  }, [time, onChange]);

  return (
    <div className="h-72 w-full flex">
      <ScrollArea className="h-full flex-1">
        <div className="w-full flex flex-col p-2">
          {hours.reverse().map((hour) => (
            <Button
              key={hour}
              size="icon"
              variant={time && time.getHours() === hour ? "default" : "ghost"}
              className="w-full shrink-0 aspect-square"
              ref={
                time && time.getHours() === hour
                  ? hourRef // Attach ref to the selected hour
                  : null
              }
              onClick={() => handleTimeChange("hour", hour.toString())}
            >
              {hour}
            </Button>
          ))}
        </div>
        <ScrollBar orientation="horizontal" className="sm:hidden" />
      </ScrollArea>
      <ScrollArea className="h-full flex-1">
        <div className="w-full flex flex-col p-2">
          {Array.from({ length: 60 }, (_, i) => i).map((minute) => (
            <Button
              key={minute}
              size="icon"
              variant={
                time && time.getMinutes() === minute ? "default" : "ghost"
              }
              className="w-full shrink-0 aspect-square"
              ref={
                time && time.getMinutes() === minute
                  ? minuteRef // Attach ref to the selected minute
                  : null
              }
              onClick={() => handleTimeChange("minute", minute.toString())}
            >
              {minute.toString().padStart(2, "0")}
            </Button>
          ))}
        </div>
        <ScrollBar orientation="horizontal" className="sm:hidden" />
      </ScrollArea>
    </div>
  );
}

export default TimePicker;
