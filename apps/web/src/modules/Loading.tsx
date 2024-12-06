import { LottieOptions, useLottie } from "lottie-react";
import popCatAnimation from "@/assets/animations/pop_cat.json";

function Loading() {
  const options: LottieOptions = {
    animationData: popCatAnimation,
    loop: true,
  };
  const { View } = useLottie(options);

  return (
    <div className="fixed w-full h-full flex justify-center items-center">
      <div className="absolute right-4 bottom-4 flex justify-end items-center gap-2">
        <div className="w-6">{View}</div>
        <p className="font-bold text-xl">Loading...</p>
      </div>
    </div>
  );
}

export default Loading;
