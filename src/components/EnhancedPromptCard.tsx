
import React, { useState } from 'react';
import { Badge } from '@/components/ui/badge';
import PromptActions from './PromptActions';
import { Dialog, DialogContent, DialogDescription, DialogHeader, DialogTitle, DialogTrigger } from "@/components/ui/dialog";
import { Button } from '@/components/ui/button';

export interface EnhancedPromptCardProps {
  id: string;
  title: string;
  description: string;
  imageUrl: string;
  tags: string[];
  trending?: boolean;
}

const MAX_DESCRIPTION_LENGTH = 120;

const EnhancedPromptCard = ({ id, title, description, imageUrl, tags, trending = false }: EnhancedPromptCardProps) => {
  const [isDialogOpen, setIsDialogOpen] = useState(false);
  const isDescriptionTruncated = description.length > MAX_DESCRIPTION_LENGTH;
  const truncatedDescription = isDescriptionTruncated 
    ? `${description.substring(0, MAX_DESCRIPTION_LENGTH)}...` 
    : description;

  return (
    <div className="glass-card rounded-xl overflow-hidden animate-fadeIn transform transition-all duration-300 hover:scale-[1.02] hover:shadow-xl flex flex-col h-full">
      {imageUrl && (
        <div className="w-full h-48 overflow-hidden">
          <img 
            src={imageUrl} 
            alt={title}
            className="w-full h-full object-cover transition-transform duration-700 hover:scale-110"
          />
        </div>
      )}
      
      <div className="p-5 flex flex-col flex-grow">
        <div className="flex-grow">
          {trending && (
            <Badge className="bg-ghibli-accent text-background mb-2">
              Trending
            </Badge>
          )}
          
          <h3 className="text-xl font-bold text-primary">{title}</h3>
          
          <div className="text-sm text-muted-foreground mt-2">
            {truncatedDescription}
            
            {isDescriptionTruncated && (
              <Dialog open={isDialogOpen} onOpenChange={setIsDialogOpen}>
                <DialogTrigger asChild>
                  <Button variant="link" className="px-0 h-auto text-xs text-ghibli-accent">
                    Show more
                  </Button>
                </DialogTrigger>
                <DialogContent className="sm:max-w-[425px]">
                  <DialogHeader>
                    <DialogTitle>{title}</DialogTitle>
                  </DialogHeader>
                  <DialogDescription className="text-muted-foreground">
                    {description}
                  </DialogDescription>
                </DialogContent>
              </Dialog>
            )}
          </div>
          
          <div className="flex flex-wrap gap-1 mt-3">
            {tags.map(tag => (
              <Badge 
                key={tag} 
                variant="outline" 
                className="text-xs bg-muted/40 text-muted-foreground border-muted/50"
              >
                {tag}
              </Badge>
            ))}
          </div>
        </div>
        
        <div className="mt-4 pt-4 border-t border-muted/20">
          <PromptActions promptId={id} promptText={description} />
        </div>
      </div>
    </div>
  );
};

export default EnhancedPromptCard;
