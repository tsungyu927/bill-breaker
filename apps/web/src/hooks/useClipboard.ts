/* eslint-disable @typescript-eslint/no-explicit-any */
import { useCallback, useRef } from "react";

export interface UseClipboardOptions {
  onSuccess?: (value: string) => void;
  onError?: (error: any) => void;
}

function useClipboard(options: UseClipboardOptions = {}) {
  const targetRef = useRef<HTMLInputElement>(null);
  const optionsRef = useRef<UseClipboardOptions>(options);
  optionsRef.current = options;

  const copyHandler = useCallback((text?: any) => {
    const opts = optionsRef.current;
    const target = targetRef.current;

    function copy(value: string) {
      navigator.clipboard
        .writeText(value)
        .then(() => {
          opts.onSuccess?.(value);
        })
        .catch((error) => {
          opts.onError?.(error);
        });
    }

    if (typeof text === "string") {
      copy(text);
    } else if (target) {
      copy(target.value);
    }
  }, []);

  return {
    ref: targetRef,
    copy: copyHandler,
  };
}

export default useClipboard;
