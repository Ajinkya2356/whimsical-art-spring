
import React from 'react';
import { Moon, Bot, Heart, TrendingUp } from 'lucide-react';
import { Button } from '@/components/ui/button';

const Header = () => {
  return (
    <header className="w-full py-4 px-6 md:px-10 flex justify-between items-center gap-4 relative z-10">
      <div className="flex items-center">
        <h1 className="text-3xl md:text-4xl bg-gradient-to-r from-ghibli-accent to-ghibli-sunset bg-clip-text text-transparent">Ghibli Prompt Garden</h1>
      </div>
      
      <div className="flex items-center gap-2">
        <Button variant="ghost" size="icon" className="rounded-full bg-muted/30">
          <Heart className="h-5 w-5 text-muted-foreground" />
        </Button>
        <Button variant="ghost" size="icon" className="rounded-full bg-muted/30">
          <TrendingUp className="h-5 w-5 text-muted-foreground" />
        </Button>
        <Button variant="ghost" size="icon" className="rounded-full bg-muted/30">
          <Bot className="h-5 w-5 text-muted-foreground" />
        </Button>
        <Button variant="ghost" size="icon" className="rounded-full bg-muted/30">
          <Moon className="h-5 w-5 text-muted-foreground" />
        </Button>
      </div>
    </header>
  );
};

export default Header;
