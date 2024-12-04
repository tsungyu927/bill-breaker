/* eslint-disable react-refresh/only-export-components */
import { Button } from "@/components/ui/button";
import {
  Card,
  CardContent,
  CardDescription,
  CardHeader,
  CardTitle,
} from "@/components/ui/card";
import {
  Dialog,
  DialogContent,
  DialogDescription,
  DialogFooter,
  DialogHeader,
  DialogTitle,
} from "@/components/ui/dialog";
import { Input } from "@/components/ui/input";
import { Label } from "@/components/ui/label";
import { Tabs, TabsList, TabsTrigger, TabsContent } from "@/components/ui/tabs";
import { RegisterAction } from "@/interface/book";
import { useState } from "react";

interface Props {
  open: boolean;
  onSubmit: ({ type, value }: { type: RegisterAction; value: string }) => void;
}

export default function RegisterDialog(props: Props) {
  const { open, onSubmit } = props;
  const [tab, setTab] = useState<RegisterAction>(RegisterAction.Create);
  const [inputValue, setInputValue] = useState<string>("");

  const handleTabChange = (tabValue: RegisterAction) => {
    setInputValue("");
    setTab(tabValue);
  };

  const handleSubmit = () => {
    onSubmit({ type: tab, value: inputValue });
  };

  return (
    <Dialog open={open}>
      <DialogContent>
        <DialogHeader>
          <DialogTitle>Create / Restore User</DialogTitle>
          <DialogDescription>
            You have never used Bill Breaker on this device. Please create a new
            user or restore from User ID.
          </DialogDescription>
        </DialogHeader>
        <Tabs
          value={tab}
          className="flex flex-col gap-2"
          onValueChange={(value) => handleTabChange(value as RegisterAction)}
        >
          <TabsList className="grid w-full grid-cols-2">
            <TabsTrigger value={RegisterAction.Create}>Create</TabsTrigger>
            <TabsTrigger value={RegisterAction.Restore} disabled>
              Restore
            </TabsTrigger>
          </TabsList>
          <TabsContent value={RegisterAction.Create}>
            <Card>
              <CardHeader>
                <CardTitle>Create new user</CardTitle>
                <CardDescription>Create a brand new user</CardDescription>
              </CardHeader>
              <CardContent>
                <Label htmlFor="name" className="text-right">
                  Name
                </Label>
                <Input
                  id="name"
                  className="col-span-3"
                  value={inputValue}
                  onChange={(e) => setInputValue(e.target.value)}
                />
              </CardContent>
            </Card>
          </TabsContent>
          <TabsContent value={RegisterAction.Restore}>
            <Card>
              <CardHeader>
                <CardTitle>Restore the user with User ID</CardTitle>
                <CardDescription>
                  You can find your User ID in your profile
                </CardDescription>
              </CardHeader>
              <CardContent>
                <Label htmlFor="userId" className="text-right">
                  User ID
                </Label>
                <Input
                  id="userId"
                  className="col-span-3"
                  value={inputValue}
                  onChange={(e) => setInputValue(e.target.value)}
                />
              </CardContent>
            </Card>
          </TabsContent>
        </Tabs>
        <DialogFooter>
          <Button onClick={handleSubmit}>Create / Restore User</Button>
        </DialogFooter>
      </DialogContent>
    </Dialog>
  );
}
