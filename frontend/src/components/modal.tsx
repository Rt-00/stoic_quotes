import { Button } from "@/components/ui/button";
import {
  Dialog,
  DialogClose,
  DialogContent,
  DialogDescription,
  DialogFooter,
  DialogHeader,
  DialogTitle,
} from "@/components/ui/dialog";

type ModalProps = {
  isOpen: boolean;
  onClose: () => void;
  quote: string;
  author: string;
  book: string;
};

export const Modal = ({ isOpen, onClose, quote, author, book }: ModalProps) => {
  return (
    <Dialog open={isOpen} onOpenChange={(open) => !open && onClose()}>
      <DialogContent className="max-h-[80vh] overflow-y-auto">
        <DialogHeader>
          <DialogTitle>Random Quote</DialogTitle>
          <DialogDescription>
            Here is randomly selected quote for you.
          </DialogDescription>
        </DialogHeader>

        <div className="overflow-y-auto max-h-[80vh] px-2">
          <blockquote className="italic text-lg mb-4">{quote}</blockquote>
          <p className="text-sm text-muted-foreground">
            - {author} ({book})
          </p>
        </div>

        <DialogFooter>
          <DialogClose asChild>
            <Button variant="secondary">Close</Button>
          </DialogClose>
        </DialogFooter>
      </DialogContent>
    </Dialog>
  );
};
