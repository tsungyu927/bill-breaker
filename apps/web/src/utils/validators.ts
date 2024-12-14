import { CostMember } from "@/interface/book";
import { ValidateResult } from "react-hook-form";

export const validateAmount = (
  costMembers: CostMember[],
  targetAmount: number
): ValidateResult | Promise<ValidateResult> => {
  const currentAmount = costMembers.reduce((acc, costMember) => {
    const amount = Number(costMember.amount);
    return acc + amount;
  }, 0);

  if (currentAmount === targetAmount) {
    return;
  }
  if (currentAmount > targetAmount) {
    return `Exceed ${currentAmount - targetAmount} dollars`;
  }
  return `Missing ${targetAmount - currentAmount} dollars`;
};
