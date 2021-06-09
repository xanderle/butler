local execute = vim.api.nvim_command
local fn = vim.fn

local install_path = fn.stdpath('data')..'/site/pack/packer/start/packer.nvim'

if fn.empty(fn.glob(install_path)) > 0 then
	  fn.system({'git', 'clone', 'https://github.com/wbthomason/packer.nvim', install_path})
	    execute 'packadd packer.nvim'
end

return require('packer').startup(function()
  use  { 'wbthomason/packer.nvim' }

  use { 'nvim-treesitter/nvim-treesitter', run = ':TSUpdate' }
  use 'nvim-treesitter/playground'
  use { 'neovim/nvim-lspconfig', config = function() 
    require 'plugins.nvim-lspconfig'
  end }

  use { 'hrsh7th/nvim-compe', config = function() 
	  require 'plugins.nvim-compe'
  end }

  -- Debugger
  use { 'mfussenegger/nvim-dap', config = function()
      require 'plugins.nvim-dap'
  end }
  -- Fuzzy finder
  use {
      'nvim-telescope/telescope.nvim',
      requires = {{'nvim-lua/popup.nvim'}, {'nvim-lua/plenary.nvim'}},
      config = function() require 'plugins.telescope' end
  }
  use {'nvim-telescope/telescope-fzf-native.nvim', run = 'make' }

  -- Vim dispatch
  use { 'tpope/vim-dispatch' }
  -- Fugitive for Git
  use { 'tpope/vim-fugitive' }
  -- Seamless tmux navigation
  use 'christoomey/vim-tmux-navigator'
  -- Color scheme
  use { 
        'marko-cerovac/material.nvim',
        config = function()
            vim.g.material_style = 'deep ocean'
            require'material'.set()
        end
    }
end)
-- https://github.com/tomaskallup/dotfiles/tree/master/nvim/lua/plugins
