import { BookMember, CostMember } from "@/interface/book";

export const calcCostMembers = (
  members: BookMember[],
  totalAmount: number
): CostMember[] => {
  const memberCount = members.length;

  if (memberCount === 0) {
    throw new Error("Members list can't be empty!");
  }

  const baseAmount = Math.floor(totalAmount / memberCount);
  const remainder = totalAmount % memberCount;

  return members.map((member, index) => ({
    userId: member.userId,
    // 會把 remainder 加在前面幾個 members 身上
    amount: baseAmount + (index < remainder ? 1 : 0),
  }));
};
