using Grpc.Core;
using Service;

namespace Service.Services;

public class CoreService : Core.CoreBase
{
    private readonly ILogger<CoreService> _logger;
    public CoreService(ILogger<CoreService> logger)
    {
        _logger = logger;
    }

    public override Task<CalculateResponse> Calculate(CalculateRequest request, ServerCallContext context)
    {
        _logger.LogInformation(request.ToString());
        
        var response = new CalculateResponse();
        response.Result = 0;

        _logger.LogInformation(response.ToString());
        return Task.FromResult(response);
    }
}
