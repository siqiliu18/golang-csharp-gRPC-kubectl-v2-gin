namespace TemplateServer.Services;

public class QHDataService : IDataService
{
    private readonly ILogger<QHDataService> _logger;
    private QhDataService.QHData.QHDataClient _client;

    public QHDataService(ILogger<QHDataService> logger, QhDataService.QHData.QHDataClient client)
    {
        _logger = logger;
        _client = client;
    }

    public string QilCall(int qilID, int qilVersion)
    {
        Console.WriteLine(">>> QilCall");
        var req = new QhDataService.QILRequest
        {
            QilID = qilID,
            QilVersion = qilVersion,
        };
        Console.WriteLine("!!! actual rpc call !!!");
        var res = _client.CallQil(req).Data;
        Console.WriteLine("@@@ res = " + res);
        return res;
    }
}