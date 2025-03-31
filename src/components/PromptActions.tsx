import React, { useState } from 'react';
import { Button } from '@/components/ui/button';
import {
  Clipboard,
  ThumbsUp,
  Bot,
  MessageSquare,
  MessageCircle,
  Code,
  Lightbulb,
  Zap,
  Check,
  ExternalLink
} from 'lucide-react';
import {
  DropdownMenu,
  DropdownMenuContent,
  DropdownMenuItem,
  DropdownMenuTrigger,
} from "@/components/ui/dropdown-menu";
import { toast } from '@/hooks/use-toast';
import supabase from '../utils/supabase';

interface PromptActionsProps {
  promptId: string;
  promptText: string;
  like_count: number;
}

const llmOptions = [
  { id: 'chatgpt', name: 'ChatGPT', icon: <MessageSquare size={16} /> },
  { id: 'claude', name: 'Claude', icon: <MessageCircle size={16} /> },
  { id: 'midjourney', name: 'Midjourney', icon: <Code size={16} /> },
  { id: 'dalle', name: 'DALL-E', icon: <Lightbulb size={16} /> },
  { id: 'stable-diffusion', name: 'Stable Diffusion', icon: <Zap size={16} /> },
];

const PromptActions = ({ promptId, promptText, like_count: initialLikeCount }: PromptActionsProps) => {
  const [liked, setLiked] = useState(false);
  const [likeCount, setLikeCount] = useState(initialLikeCount);

  const copyToClipboard = () => {
    navigator.clipboard.writeText(promptText);
    toast({
      title: "Copied!",
      description: "Prompt copied to clipboard",
    });
  };

  const toggleLike = async () => {
    const newLikedState = !liked;
    const newLikeCount = newLikedState ? likeCount + 1 : likeCount - 1;

    setLiked(newLikedState);
    setLikeCount(newLikeCount);

    toast({
      title: newLikedState ? "Added to favorites" : "Removed from favorites",
      description: newLikedState ? "Prompt added to your favorites" : "Prompt removed from your favorites",
    });

    // Optimistically update the UI, then update Supabase in the background
    try {
      const { error } = await supabase.rpc('increment_like_count', {
        prompt_id: promptId
      });

      if (error) {
        console.error("Error updating like count:", error);
        // Revert the local state if the update fails
        setLiked(liked);
        setLikeCount(likeCount);
        toast({
          title: "Error",
          description: "Failed to update like count. Please try again.",
        });
      }
    } catch (error) {
      console.error("Error updating like count:", error);
      // Revert the local state if the update fails
      setLiked(liked);
      setLikeCount(likeCount);
      toast({
        title: "Error",
        description: "Failed to update like count. Please try again.",
      });
    }
  };

  const openLLM = (llmId: string) => {
    // In a real app, this would redirect to the appropriate LLM with the prompt
    let url = '';

    switch (llmId) {
      case 'chatgpt':
        url = `https://chat.openai.com/new?prompt=${encodeURIComponent(promptText)}`;
        break;
      case 'claude':
        url = `https://claude.ai?prompt=${encodeURIComponent(promptText)}`;
        break;
      case 'midjourney':
        url = `https://www.midjourney.com?prompt=${encodeURIComponent(promptText)}`;
        break;
      case 'dalle':
        url = `https://labs.openai.com?prompt=${encodeURIComponent(promptText)}`;
        break;
      case 'stable-diffusion':
        url = `https://stability.ai?prompt=${encodeURIComponent(promptText)}`;
        break;
      default:
        url = '';
    }

    if (url) {
      window.open(url, '_blank');
    }
  };

  return (
    <div className="flex items-center gap-2">
      <Button
        variant="outline"
        size="sm"
        onClick={copyToClipboard}
        className="text-muted-foreground hover:text-primary border-muted"
      >
        <Clipboard size={16} className="mr-1" />
        Copy
      </Button>

      <Button
        variant="outline"
        size="sm"
        onClick={toggleLike}
        className={`border-muted ${liked ? 'text-red-400 hover:text-red-500' : 'text-muted-foreground hover:text-primary'}`}
      >
        <ThumbsUp size={16} className="mr-1" fill={liked ? 'currentColor' : 'none'} />
        {likeCount}
      </Button>

      <DropdownMenu>
        <DropdownMenuTrigger asChild>
          <Button
            variant="outline"
            size="sm"
            className="text-muted-foreground hover:text-primary border-muted"
          >
            <Bot size={16} className="mr-1" />
            Use LLM
          </Button>
        </DropdownMenuTrigger>

        <DropdownMenuContent align="end" className="bg-card w-48">
          {llmOptions.map(option => (
            <DropdownMenuItem
              key={option.id}
              onClick={() => openLLM(option.id)}
              className="flex items-center gap-2 cursor-pointer"
            >
              {option.icon}
              <span>{option.name}</span>
              <ExternalLink size={12} className="ml-auto opacity-50" />
            </DropdownMenuItem>
          ))}
        </DropdownMenuContent>
      </DropdownMenu>
    </div>
  );
};

export default PromptActions;
