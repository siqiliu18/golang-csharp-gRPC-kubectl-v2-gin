using Grpc.Core;
using TemplateServer;

namespace TemplateServer.Services;

public class TemplateService : TemplateMaker.TemplateMakerBase
{
    private readonly ILogger<TemplateService> _logger;
    private IDataService _dataService; // after creating IDataService interface, need to use it here
    public TemplateService(ILogger<TemplateService> logger, IDataService dataService)
    {
        _logger = logger;
        _dataService = dataService;
    }

    public override Task<TemplateResponse> RetrieveTemplate(TemplateRequest request, ServerCallContext context)
    {
        Console.WriteLine(" Call QIL before: ");
        var res = _dataService.QilCall(6, 1);
        Console.WriteLine(" QilCall works? " + res);

        return Task.FromResult(new TemplateResponse
        {
            Documenturi = request.Templatename + "." + request.Templatetype + " - " + request.Cbedna + " : " + res
        });
    }
}
