import { FormField, FormItem } from "@/components/ui/form";
import { Input } from "@/components/ui/input";
import {
  Table,
  TableBody,
  TableCell,
  TableHead,
  TableHeader,
  TableRow,
} from "@/components/ui/table";
import { BookMember, CreateCostForm } from "@/interface/book";
import { validateAmount } from "@/utils/validators";
import { useFieldArray, useFormContext, useWatch } from "react-hook-form";

interface Props {
  name: "payers" | "sharers";
  members: BookMember[];
}

function MemberSelectList(props: Props) {
  const { name, members } = props;
  const {
    control,
    trigger,
    formState: { errors },
  } = useFormContext<CreateCostForm>();
  const { fields } = useFieldArray({
    control,
    name,
    rules: {
      validate: (list) => validateAmount(list, watchAmount),
    },
  });

  const watchAmount = useWatch({ control, name: "amount" });
  const errorMessage = errors[name]?.root?.message;

  const handleTriggerValidator = () => {
    trigger(name);
  };

  return (
    <div>
      <Table>
        <TableHeader>
          <TableRow>
            <TableHead>Name</TableHead>
            <TableHead>Amount</TableHead>
          </TableRow>
        </TableHeader>
        <TableBody>
          {fields.map((field, index) => {
            const member = members.find(
              (bookMember) => bookMember.userId === field.userId
            );
            return (
              <TableRow key={field.id}>
                <TableCell>{member?.userName}</TableCell>
                <TableCell>
                  <FormField
                    control={control}
                    name={`${name}.${index}.amount`}
                    render={({ field: { value, onChange } }) => (
                      <FormItem>
                        <Input
                          type="number"
                          value={value}
                          onBlur={handleTriggerValidator}
                          onChange={(e) => {
                            if (e.target.value === "") {
                              return onChange("");
                            }
                            onChange(Number(e.target.value));
                          }}
                        />
                      </FormItem>
                    )}
                  />
                </TableCell>
              </TableRow>
            );
          })}
        </TableBody>
      </Table>
      {errorMessage && (
        <p className="text-[0.8rem] font-medium text-destructive">
          {errorMessage}
        </p>
      )}
    </div>
  );
}

export default MemberSelectList;
