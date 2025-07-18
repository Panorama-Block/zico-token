import * as React from "react"
import { cn } from "../../lib/utils"

const buttonVariants = {
  default: "bg-gradient-to-r from-purple-600 to-pink-600 text-white hover:from-purple-700 hover:to-pink-700",
  secondary: "bg-gray-100 text-gray-900 hover:bg-gray-200 border border-gray-200",
  outline: "border border-gray-300 bg-transparent hover:bg-gray-100 text-gray-700",
  ghost: "hover:bg-gray-100 text-gray-700",
  destructive: "bg-red-500 text-white hover:bg-red-600",
}

const buttonSizes = {
  default: "h-10 px-4 py-2",
  sm: "h-9 rounded-md px-3",
  lg: "h-11 rounded-md px-8",
  icon: "h-10 w-10",
}

const Button = React.forwardRef(({ 
  className, 
  variant = "default", 
  size = "default", 
  ...props 
}, ref) => {
  return (
    <button
      className={cn(
        "inline-flex items-center justify-center rounded-lg text-sm font-medium transition-all duration-200 focus-visible:outline-none focus-visible:ring-2 focus-visible:ring-ring focus-visible:ring-offset-2 disabled:pointer-events-none disabled:opacity-50",
        buttonVariants[variant],
        buttonSizes[size],
        className
      )}
      ref={ref}
      {...props}
    />
  )
})
Button.displayName = "Button"

export { Button } 