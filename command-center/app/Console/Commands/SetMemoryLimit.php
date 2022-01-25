<?php

namespace App\Console\Commands;

use Illuminate\Console\Command;
use GuzzleHttp\Client;

class SetMemoryLimit extends Command
{
    /**
     * The name and signature of the console command.
     *
     * @var string
     */
    protected $signature = 'helper:setmemory {limit}';
    /**
     * The console command description.
     *
     * @var string
     */
    protected $description = 'Pass memory limit to agent';

    /**
     * Create a new command instance.
     *
     * @return void
     */
    public function __construct()
    {
        parent::__construct();
    }

    /**
     * Execute the console command.
     *
     * @return int
     */
    public function handle()
    {
        $limit = $this->argument('limit');
        $response = $client = new Client([
            // Base URI is used with relative requests
            'base_uri' => 'http://localhost:8080/memory-limit/' . $limit,
            // You can set any number of default request options.
            'timeout'  => 2.0,
        ]);

        $response = $response->getBody()->getContents();

        print_r($response);
    }
}
